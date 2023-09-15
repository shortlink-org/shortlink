## Kube Secret Reader

This Node.js script fetches a secret from a Kubernetes cluster and writes it to a .env file.

### Prerequisites

- Node.js
- npm (Node package manager)
- A Kubernetes cluster where you have the necessary permissions to get secrets.

### Usage

Run the script with the following command:

```bash
npm run start -- --namespace <namespace> --secret <secret-name> --key <key> --envKey <envKey>
```

Replace `<namespace>`, `<secret-name>`, `<key>` and `<envKey>` with your values.

- **namespace**: the Kubernetes namespace where your secret is stored.
- **secret**: the name of your secret.
- **key**: the key of the value you want to get from the secret.
- **envKey**: the key you want to use when writing the secret's value to the .env file.
- **envPath** _(optional)_: the path to your .env file. 

This script assumes that the .env file is located in the same directory where the script is run.
