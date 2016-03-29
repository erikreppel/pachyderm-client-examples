package main

import (
    "log"
    "github.com/pachyderm/pachyderm/src/client"
    "github.com/pachyderm/pachyderm/src/client/pfs"
)

func main() {
    // Connect to Pachyderm
    // That IP is probably not constant, run `kubectl get all`
    // and see what IP pachd is running at
    apiClient, err := client.NewFromAddress("10.0.0.115:650")
    if err != nil{
        log.Fatal("Error:", err)
    }
    
    log.Println("Established client")    
    
    // Get a list of our repos
    repos, err := pfs.ListRepo(apiClient)
    if err != nil {
        log.Println(err)
    }
    
    for _, info := range repos {
        log.Println("Name:", info.Repo.Name,
                    "Created:", info.Created,
                    "Size", info.SizeBytes)
    }
    // fmt.Println(repoInfo) // this work too
    
    // Lets get a file from a commit in one of the repos
    repo := repos[0].Repo.Name
    
    repoCommits, err := pfs.ListCommit(apiClient,
                                       []string{repo},
                                       []string{},
                                       false)
    if err != nil { log.Fatal(err) }
    
    log.Println("Commits in", repo, ":", repoCommits)
    
    commitID := repoCommits[0].Commit.ID
    
    // Lets get some files
    log.Println("Getting files from commit", commitID, "in", repo)
    
    // First we have to inspect a file to get the shard, we need a list of files
    // Need to figure out how to get a shard
    // fileInfo, err := pfs.ListFile(apiClient, repo, commitID, "./", commitID, &shard)
    
}