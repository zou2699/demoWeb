def label = "mypod-${UUID.randomUUID().toString()}"

podTemplate(label: label,
containers: [
    containerTemplate(name: 'docker', image: 'docker:18.06.2', command: 'cat', ttyEnabled: true)],
volumes: [
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock')]
)  {

    node(label) {
        def ImageName = 'zouhl/demo-app'

        stage('Get a Golang project'){
            git url: 'https://github.com/zou2699/demoweb.git'
            container('docker') {
                stage('Build a Go project and docker image')
                sh """
                docker build -t ${ImageName}:${BUILD_NUMBER} .
                """

                stage('push image') {
                withDockerRegistry(credentialsId: 'dockerHub') {
                    sh """
                    docker push ${ImageName}:${BUILD_NUMBER}
                    """
                    }
                }

                stage('Remove the local docker image') {
                    sh "docker rmi ${ImageName}:${BUILD_NUMBER}"
                }
            }
        }
        stage('deploy to k8s') {
           kubernetesDeploy(
               kubeconfigId: 'kubeconfig',
               configs: 'deploy.yaml',
               enableConfigSubstitution: true,
               )
        }
        stage('create or update ingress') {
            input 'Ready to go?'
            kubernetesDeploy(
               kubeconfigId: 'kubeconfig',
               configs: 'app-ingress.yaml',
               enableConfigSubstitution: true,
               )
        }
}
}