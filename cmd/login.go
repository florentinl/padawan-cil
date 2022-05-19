package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to padawan",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Go to https://padawan.kube.test.viarezo.fr/login")
			fmt.Println("Then paste the value of the _forward_auth cookie in the command line")
			fmt.Println("")
			fmt.Println("Example:")
			fmt.Println("$ padawan login \"<cookie>\"")
			fmt.Println("")
		} else {
			cookie := args[0]
			fmt.Println("Logging in...")
			req, err := http.NewRequest("GET", apiUri()+"/login", nil)
			if err != nil {
				fmt.Println("Error while logging in")
				os.Exit(1)
			}
			req.AddCookie(&http.Cookie{Name: "_forward_auth", Value: cookie})
			response, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println("Error while logging in")
				os.Exit(1)
			}
			if response.StatusCode != 200 {
				fmt.Println("Error while logging in")
				os.Exit(1)
			}
			setToken(cookie)
			fmt.Println("Logged in")
			fmt.Println("")
			fmt.Println("You can now use the padawan command")
			fmt.Println("")
			fmt.Println("Example:")
			fmt.Println("$ padawan")
			fmt.Println("")
		}
	},
}
