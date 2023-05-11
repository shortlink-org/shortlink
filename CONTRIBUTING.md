# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change. 

Please note we have a code of conduct, please follow it in all your interactions with the project.

### Getting Started

```
# Install GIT sub-repository
git submodule update --init --recursive

# Setting up the environment
cp .env.example .env
```

#### Different launch options 

<details><summary>DETAILS</summary>
<p>

##### docker compose

<details><summary>DETAILS</summary>
<p>

###### For run
```
# Run all stack. WARNING: This command run a lot of containers
make run

# Run only the containers that are necessary scope of the project
make dev
```

###### For down

```
make down
```

</p>
</details>

##### Kubernetes (1.26+)

<details><summary>DETAILS</summary>
<p>

###### For run
```
make minikube-up
make helm-shortlink-up
```

###### For down
```
make minikube-down
```

</p>
</details>

##### Skaffold [(link)](https://skaffold.dev/)

<details><summary>DETAILS</summary>
<p>

###### For run
```
make skaffold-init
make skaffold-up
```

###### For down
```
make skaffold-down
```

###### Debug mode
```
make skaffold-debug
```

</p>
</details>

</p>
</details>
