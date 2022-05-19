package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Manage images",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Manage Padawan images")
	},
}

var imgLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List images",
	Run: func(cmd *cobra.Command, args []string) {
		padawan_token := getToken()
		req, err := http.NewRequest("GET", apiUri()+"/images", nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.AddCookie(&http.Cookie{Name: "_forward_auth", Value: padawan_token})
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer response.Body.Close()
		if response.StatusCode == 401 {
			displayLoginMessage()
		}
		if response.StatusCode != 200 {
			fmt.Println("Error while listing images")
			os.Exit(1)
		}
		b := make([]byte, 1000)
		_, err = response.Body.Read(b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	},
}

var imgAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an image",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("You must specify an image name, a repository and a shell")
			fmt.Println("Example:")
			fmt.Println("$ padawan img add <image_name> <repository> <shell>")
			os.Exit(1)
		}
		imageRequest := []byte(`{
			"image_name": "` + args[0] + `",
			"repository": "` + args[1] + `",
			"shell": "` + args[2] + `"
		}`)
		padawan_token := getToken()
		req, err := http.NewRequest("POST", apiUri()+"/images", bytes.NewBuffer(imageRequest))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.AddCookie(&http.Cookie{Name: "_forward_auth", Value: padawan_token})
		req.Header.Set("Content-Type", "application/json")
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if response.StatusCode == 401 {
			displayLoginMessage()
		}
		if response.StatusCode != 200 {
			fmt.Println("Error while adding image")
			os.Exit(1)
		}
		fmt.Println("Image added")
	},
}

var imgDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete an image",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("You must specify an image name")
			fmt.Println("Example:")
			fmt.Println("$ padawan img del <image_name>")
			os.Exit(1)
		}
		padawan_token := getToken()
		req, err := http.NewRequest("DELETE", apiUri()+"/images/"+args[0], nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.AddCookie(&http.Cookie{Name: "_forward_auth", Value: padawan_token})
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if response.StatusCode == 401 {
			displayLoginMessage()
		}
		if response.StatusCode != 200 {
			fmt.Println("Error while deleting image")
			os.Exit(1)
		}
		fmt.Println("Image deleted")
	},
}

var imgGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an image",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("You must specify an image name")
			fmt.Println("Example:")
			fmt.Println("$ padawan img get <image_name>")
			os.Exit(1)
		}
		padawan_token := getToken()
		req, err := http.NewRequest("GET", apiUri()+"/images/"+args[0], nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.AddCookie(&http.Cookie{Name: "_forward_auth", Value: padawan_token})
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if response.StatusCode == 401 {
			displayLoginMessage()
		}
		if response.StatusCode != 200 {
			fmt.Println("Error while getting image")
			os.Exit(1)
		}
		b := make([]byte, 1000)
		_, err = response.Body.Read(b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(b))
	},
}

var imgSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an image",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println("You must specify an image name, a repository and a shell")
			fmt.Println("Example:")
			fmt.Println("$ padawan img set <image_name> <repository> <shell>")
			os.Exit(1)
		}
		imageRequest := []byte(`{
			"image_name": "` + args[0] + `",
			"repository": "` + args[1] + `",
			"shell": "` + args[2] + `"
		}`)
		padawan_token := getToken()
		req, err := http.NewRequest("PUT", apiUri()+"/images/"+args[0], bytes.NewBuffer(imageRequest))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.AddCookie(&http.Cookie{Name: "_forward_auth", Value: padawan_token})
		req.Header.Set("Content-Type", "application/json")
		response, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if response.StatusCode == 401 {
			displayLoginMessage()
		}
		if response.StatusCode != 200 {
			fmt.Println("Error while setting image")
			os.Exit(1)
		}
		fmt.Println("Image set")
	},
}
