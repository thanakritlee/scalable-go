pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'cd entryserver && go build'
            }
        }
    }
}