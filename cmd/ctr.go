package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Container struct {
	Username  string `json:"username"`
	ImageName string `json:"image_name"`
	Port      int    `json:"port"`
}

var ctrCmd = &cobra.Command{
	Use:   "ctr",
	Short: "Manage containers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Manage containers")
	},
}

var ctrLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List containers",
	Run: func(cmd *cobra.Command, args []string) {
		padawan_token := getToken()
		req, err := http.NewRequest("GET", apiUri()+"/containers", nil)
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
			fmt.Println("Error while listing containers")
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

var ctrRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a container",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("You must provide an image name")
			fmt.Println("Example:")
			fmt.Println("$ padawan ctr run <image>")
			os.Exit(1)
		}
		containerRequest := []byte(fmt.Sprintf(`{"image_name":"%s", "password":"%s"}`,
			args[0],
			getPassword()))

		padawan_token := getToken()
		req, err := http.NewRequest("POST", apiUri()+"/container/me", bytes.NewBuffer(containerRequest))
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
			fmt.Println("Error while running container")
			os.Exit(1)
		}
		var container Container
		json.NewDecoder(response.Body).Decode(&container)
		fmt.Println("Container successfully created from image: " + container.ImageName)
		fmt.Println("Connect to your padawan container with ssh:")
		fmt.Printf("$ ssh %s@%s -p %d\n", container.Username, apiFQDN(), container.Port)
	},
}

var ctrDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete containers",
	Run: func(cmd *cobra.Command, args []string) {
		var route string
		if len(args) == 0 {
			route = apiUri() + "/container/me"
		} else {
			if len(args) == 1 {
				route = apiUri() + "/container/" + args[0]
			} else {
				fmt.Println("You must provide a container name or avoid the argument")
				fmt.Println("Example:")
				fmt.Println("$ padawan ctr del <container>")
				fmt.Println("$ padawan ctr del")
				os.Exit(1)
			}
		}

		padawan_token := getToken()
		req, err := http.NewRequest("DELETE", route, nil)
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
			fmt.Println("Error while deleting container")
			os.Exit(1)
		}
		fmt.Println("Container deleted")
	},
}

var ctrGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get containers",
	Run: func(cmd *cobra.Command, args []string) {
		var route string
		if len(args) == 0 {
			route = apiUri() + "/container/me"
		} else {
			if len(args) == 1 {
				route = apiUri() + "/container/" + args[0]
			} else {
				fmt.Println("You must provide a container name or avoid the argument")
				fmt.Println("Example:")
				fmt.Println("$ padawan ctr get <container>")
				fmt.Println("$ padawan ctr get")
				os.Exit(1)
			}
		}

		padawan_token := getToken()
		req, err := http.NewRequest("GET", route, nil)
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
			if len(args) == 0 {
				fmt.Println("You do not have a container yet !")
				fmt.Println("Create one now with:")
				fmt.Println("$ padawan ctr run <image>")
			} else {
				fmt.Println("No container with name " + args[0])
			}

			os.Exit(1)
		}
		var container Container
		json.NewDecoder(response.Body).Decode(&container)
		fmt.Println("Container is running from image: " + container.ImageName)
		fmt.Println("Connect to your padawan container with ssh:")
		fmt.Printf("$ ssh %s@%s -p %d\n", container.Username, apiFQDN(), container.Port)
	},
}
