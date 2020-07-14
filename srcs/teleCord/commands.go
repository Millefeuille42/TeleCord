package teleCord

import (
	"encoding/json"
	"fmt"
	"github.com/Millefeuille42/TeleCord/definitions"
	"github.com/Millefeuille42/TeleCord/teleCord/routers"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Register(origin string, message definitions.MessageStruct) error {
	path := fmt.Sprintf("../data/%s/%d.json", origin, message.SenderID)
	arg := strings.Split(message.MessageContent, "-")
	if len(arg) <= 3 {
		_ = routers.SendingRouter(origin, "Invalid number of arguments", message.SenderID, nil)
		return nil
	}

	arg[1] = strings.TrimSpace(arg[1])
	arg[2] = strings.TrimSpace(arg[2])
	arg[2] = strings.ToLower(arg[2])
	arg[3] = strings.TrimSpace(arg[3])
	dest, _ := strconv.Atoi(arg[1])
	contact := definitions.Contact{
		FromID:     message.SenderID,
		ToID:       dest,
		ToPlatform: arg[2],
		ToName:     arg[3],
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	jsonData, err := json.MarshalIndent(contact, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	err = routers.SendingRouter(origin, fmt.Sprintf("You are now talking to %s (%s) on %s", arg[3], arg[1], arg[2]),
		message.SenderID, nil)
	if err != nil {
		return err
	}
	return nil
}

func GetDest(origin string, message definitions.MessageStruct) error {
	var path = fmt.Sprintf("./data/%s/%d.json", origin, message.SenderID)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = routers.SendingRouter(origin, "You must register using /dest -ID -Platform -Nickname",
			message.SenderID, nil)
		return nil
	}
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fileDataJson := definitions.Contact{}
	err = json.Unmarshal(fileData, &fileDataJson)
	if err != nil {
		return err
	}

	response := fmt.Sprintf("Currently talking to %s (%d) on %s",
		fileDataJson.ToName, fileDataJson.ToID, fileDataJson.ToPlatform)
	err = routers.SendingRouter(origin, response, message.SenderID, nil)
	if err != nil {
		return err
	}
	return err
}

func TransmitMessage(origin string, message definitions.MessageStruct) error {
	var path = fmt.Sprintf("../data/%s/%d.json", origin, message.SenderID)

	message.MessageContent = fmt.Sprintf("FROM: %s (%d - %s)\n\t%s", message.SenderName, message.SenderID,
		origin, message.MessageContent)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = routers.SendingRouter(origin, "You must register using /dest -ID -Platform -Nickname",
			message.SenderID, nil)
		return nil
	}
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fileDataJson := definitions.Contact{}
	err = json.Unmarshal(fileData, &fileDataJson)
	if err != nil {
		return err
	}
	if fileDataJson.ToPlatform == "discord" {
		message.MessageContent = fmt.Sprintf("```%s```", message.MessageContent)
	}
	err = routers.SendingRouter(fileDataJson.ToPlatform, message.MessageContent, fileDataJson.ToID, message.AttachmentsLinks)
	if err != nil {
		return err
	}
	return nil
}

func HandleMessage(origin string) {
	message, err := routers.ParsingRouter(origin)
	if err != nil {
		_ = routers.SendingRouter(origin, err.Error(), message.SenderID, nil)
		return
	}

	if strings.HasPrefix(message.MessageContent, "/") {
		var err error = nil

		if strings.HasPrefix(message.MessageContent, "/dest") {
			err = Register(origin, message)
		}
		if strings.HasPrefix(message.MessageContent, "/myID") {
			err = routers.SendingRouter(origin, fmt.Sprintf("Your ID:\n\t%d", message.SenderID), message.SenderID, nil)
		}
		if strings.HasPrefix(message.MessageContent, "/myDest") {
			err = GetDest(origin, message)
		}
		if err != nil {
			_ = routers.SendingRouter(origin, err.Error(), message.SenderID, nil)
		}
		return
	}

	err = TransmitMessage(origin, message)
	if err != nil {
		_ = routers.SendingRouter(origin, err.Error(), message.SenderID, nil)
	}
}
