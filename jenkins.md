
# jenkins

> automation tool for software development process like building, testing, and deploying applications

- extensible
  - plugins enable integration with tools like git, docker, kubernetes
- cross-platform
  - runs on windows, linux, macos
- used for CI/CD


- job: single task or series of steps
- pipeline: sequence of automated steps that handle building, testing and deploying code
  - provides visualization of progress and status
  - supports both declarative and scripted syntax
  - defined in a Jenkinsfile
  - uses language similar to groovy

example of Jenkinsfile

```groovy
pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo 'Building...'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing...'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying...'
            }
        }
    }
}
```

- node (master): main jenkins server that orchestrates jobs and manages UI
- agent (slave): machines where jenkins executes tasks

- triggers: defines how a job or a pipeline starts
  - examples:
    - manual: Click "Build Now"
    - SCM polling: checks for changes in a repository
    - webhook: triggered by external tools like GitHub after a commit
    - cron: scheduled execution at specific times

## setup

can be installed manually or with docker in any machine

### AWS EC2 Instance Setup

1. Create an Ubuntu EC2 instance on AWS
2. Configure Security Groups:
   - Allow TCP port 8080 for Jenkins
   - Allow SSH port 22 for remote access
   - Configure other necessary ports based on your application

### Jenkins Installation

1. Update package index:

```bash
sudo apt update
```

2. Install Java (Jenkins requirement):

```bash
sudo apt install openjdk-11-jdk
```

3. Install Jenkins:

```bash
curl -fsSL https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key | sudo tee \
  /usr/share/keyrings/jenkins-keyring.asc > /dev/null
echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
  https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
  /etc/apt/sources.list.d/jenkins.list > /dev/null
sudo apt update
sudo apt install jenkins
```

### Accessing Jenkins

1. Get your EC2 instance's public IPv4 DNS
2. Access Jenkins via web browser: `http://<your-instance-public-dns>:8080`
3. Retrieve initial admin password:

```bash
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
```

### Jenkins Configuration

1. Create admin account
2. Install suggested plugins
3. Configure system settings as needed

### Creating a Pipeline

1. Click "New Item" on Jenkins dashboard
2. Choose project type:
   - Freestyle: Simple, single tasks
   - Pipeline: Complex, multi-stage builds
   - Multi-branch Pipeline: Automatic branch detection

### GitHub Integration

1. Generate SSH key for Jenkins:

```bash
ssh-keygen -t ed25519 -C "jenkins@example.com"
```

2. Add public key to GitHub repository settings
3. Configure Jenkins credentials:
   - Add SSH key or
   - Use Personal Access Token (PAT)

### Pipeline Configuration

1. Select "Pipeline script from SCM" if using Jenkinsfile
2. Enter repository URL
3. Configure credentials
4. Specify branch to build (main/master)
5. Set path to Jenkinsfile if different from default

### Best Practices

1. Use declarative pipeline syntax
2. Store sensitive information in Jenkins credentials
3. Implement proper error handling
4. Set up notifications for build status
5. Regular backup of Jenkins configuration

### Troubleshooting

- Verify correct branch name (main vs master)
- Check network connectivity
- Validate credentials
- Review Jenkins logs for errors
- ENSURE PROPER PERMISSIONS ON FILES/DIRECTORIES
