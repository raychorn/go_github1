package main

import (
	"context"
	"fmt"
	//"reflect"
	"log"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func fetchOrganizations(username string) ([]*github.Organization, error) {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

func main() {
	fmt.Println("Github.com")

		ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "23b73191805dbc36047ad8db6311729c32442aca"},
	)
	tc := oauth2.NewClient(ctx, ts)

	//fmt.Printf("%s\n", tc)

	client := github.NewClient(tc)

	user, _, err := client.Users.Get(ctx, "")

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	fmt.Printf("\n%v -> %s\n", github.Stringify(user), user.GetLogin())	

	organizations, err := fetchOrganizations(user.GetLogin())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("BEGIN ORGS:")
	for i, organization := range organizations {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}
	fmt.Println("END ORGS !!!")
	fmt.Println()

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "raychorn", nil)

	if (err != nil) {
		log.Fatal(err)
		panic("Oops.")
	}

	count := 0
	for ii,repo := range repos {
		fmt.Println("BEGIN:")
		fmt.Printf("\tRepo #%d\n", ii)
		fmt.Printf("\t\tID --> %d\n", repo.ID)
		if (repo.Name != nil) {
			fmt.Printf("\t\tName --> %s\n", *repo.Name)
		}
		if (repo.Description != nil) {
			fmt.Printf("\t\tDescription --> %s\n", *repo.Description)
		}
		if (repo.License != nil) {
			fmt.Printf("\t\tLicense --> %s\n", repo.License)
		}
		if (repo.Private != nil) {
			fmt.Printf("\t\tPrivate --> %t\n", *repo.Private)
		}
		if (repo.HasIssues != nil) {
			fmt.Printf("\t\tHasIssues --> %t\n", *repo.HasIssues)
		}
		fmt.Println("END!!!")
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println()

		count++
	}
	fmt.Printf("There are %d repos.", count)
	
}