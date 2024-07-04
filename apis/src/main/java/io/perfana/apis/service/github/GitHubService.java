package io.perfana.apis.service.github;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.reactive.function.client.WebClient;

import java.util.List;

@Service
public class GitHubService {

    private final WebClient webClient;

    @Autowired
    public GitHubService(WebClient webClient) {
        this.webClient = webClient;
    }

    public List<Object> getRepositories() {
        return webClient
                .get()
                .uri("https://api.github.com/user/repos")
                .retrieve()
                .bodyToMono(List.class)
                .block();
    }

    public List<Object> getBranches(String owner, String repo) {
        String url = String.format("https://api.github.com/repos/%s/%s/branches", owner, repo);
        return webClient
                .get()
                .uri(url)
                .retrieve()
                .bodyToMono(List.class)
                .block();
    }

    public List<Object> getCommits(String owner, String repo, String branch) {
        String url = String.format("https://api.github.com/repos/%s/%s/commits?sha=%s", owner, repo, branch);
        return webClient
                .get()
                .uri(url)
                .retrieve()
                .bodyToMono(List.class)
                .block();
    }
}