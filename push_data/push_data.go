package main

import (
    "log"
    "github.com/pachyderm/pachyderm/src/client"
    "github.com/pachyderm/pachyderm/src/client/pfs"
)

func main() {
    apiClient, err := client.NewFromAddress("10.0.0.115:650")
    if err != nil { log.Fatal(err) }
    
    log.Println("Established connection to Pachyderm")
    
    repoName := "testRepo"
    log.Println("Creating a new repo named", repoName)
    
    pfs.CreateRepo(apiClient, repoName)
    log.Println("Successfully created the repo", repoName)
    
    repoInfo, err := pfs.InspectRepo(apiClient, repoName)
    if err != nil { log.Fatal(err) }
    log.Println(repoInfo)
    
}

