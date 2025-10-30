const { exec } = require('child_process');
const util = require('util');
const yargs = require('yargs/yargs');
const { hideBin } = require('yargs/helpers');
const fs = require('fs');
const os = require('os');

const execPromisified = util.promisify(exec);

async function main() {
  const argv = yargs(hideBin(process.argv)).options({
    'namespace': { type: 'string', demandOption: true, describe: 'Kubernetes namespace' },
    'secret': { type: 'string', demandOption: true, describe: 'Secret name' },
    'key': { type: 'string', demandOption: true, describe: 'Key inside the secret' },
    'envKey': { type: 'string', demandOption: true, describe: 'Key for the .env file' },
    'envPath': { type: 'string', default: '.env', describe: 'Path to the .env file' }
  }).argv;

  const { namespace, secret: secretName, key, envKey, envPath } = argv;

  try {
    const { stdout } = await execPromisified(`kubectl -n ${namespace} get secret ${secretName} -o json`);
    const secret = JSON.parse(stdout);
    const keyValueBase64 = secret.data[key];
    const keyValueDecoded = Buffer.from(keyValueBase64, 'base64').toString().replace('.svc', '');

    let envConfig = fs.readFileSync(envPath, 'utf8').split(os.EOL);
    let newConfig = envConfig.map(line => {
      if (line.startsWith(envKey)) {
        return `${envKey}=${keyValueDecoded}`
      }
      return line;
    });

    if (!newConfig.includes(`${envKey}=${keyValueDecoded}`)) {
      newConfig.push(`${envKey}=${keyValueDecoded}`);
    }

    fs.writeFileSync(envPath, newConfig.join(os.EOL));
    console.log(`Written to ${envPath}: ${envKey}=${keyValueDecoded}`);
  } catch (error) {
    console.error(`Error: ${error}`);
  }
}

main();
