package main

import (
    "log"
    "github.com/pachyderm/pachyderm/src/client"
    "github.com/pachyderm/pachyderm/src/client/pfs"
)

func shard() *pfs.Shard {
    var fileNumber int
    var fileModulus int
    var blockNumber int
    var blockModulus int
    return &pfs.Shard {
        FileNumber:   uint64(fileNumber),
        FileModulus:  uint64(fileModulus),
        BlockNumber:  uint64(blockNumber),
        BlockModulus: uint64(blockModulus),
    }
}

func main() {

    apiClient, err := client.NewFromAddress("10.0.0.244:650")
    if err != nil { log.Fatal(err) }
    
    log.Println("Established connection to Pachyderm")
    
    repoName := "testRepo"
    log.Println("Creating a new repo named", repoName)
    
    pfs.CreateRepo(apiClient, repoName)
    log.Println("Successfully created the repo", repoName)
    
    repoInfo, err := pfs.InspectRepo(apiClient, repoName)
    if err != nil { log.Fatal(err) }
    log.Println(repoInfo)
    
    log.Println("Starting a commit")
    branch := ""
    var parentCommitID string
    commit, err := pfs.StartCommit(apiClient, repoName, parentCommitID, branch)
    if err != nil {
        log.Fatal("Failed to start commit")

    }
    
}

