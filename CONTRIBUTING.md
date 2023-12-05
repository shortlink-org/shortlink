# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change. 

> [!WARNING]  
> Please note we have a code of conduct, please follow it in all your interactions with the project.

### Initial Setup

Firstly, you need to set up your environment:

```
# Install GIT sub-repository
git submodule update --init --recursive

# Setting up the environment
cp .env.example .env
```

#### Launch Options 

Depending on your preference, you can launch the project using either Docker Compose, Minikube/Kubernetes (1.28+), 
or Skaffold. Here are the details for each method:

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

##### Minikube/Kubernetes (1.28+)

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

</p></details>

</p></details>

### Branch Naming Convention

When undertaking a new task, it's generally a good practice to create a separate branch. This approach prevents excessive changes in a single commit. However, in some cases where tasks are closely related, combining them might be more beneficial.

> [!TIP]
> The naming conventions for the branches are as follows:
>
> - For new features, use the **feature/new-feature-description** pattern. 
> Ensure that the description succinctly summarizes the changes.
> - For bug fixes, use the **fix/** prefix followed by a brief summary of the fix.

#### Code Workflow

We adhere to the `green master` methodology, aiming to integrate changes into the main branch as swiftly as possible
for deployment to a testing environment, and ideally, to production.

After implementing a task in a separate branch, create a Merge Request (MR).
The MR description should ideally include any additional context or details about the implementation.

Continuous Integration (CI) checks must pass. It's essential to write enough tests to ensure confidence in the changes,
including, at minimum, tests for primary positive scenarios. After an MR has been approved, if the reviewer
hasn't merged it, feel free to proceed with merging it yourself.
