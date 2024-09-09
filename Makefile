.PHONY: jenkins-build
jenkins-build: ;

.PHONY: jenkins-run
jenkins-run:
	java -Djenkins.model.Jenkins.crumbIssuerProxyCompatibility=true -Dhudson.model.DirectoryBrowserSupport.CSP="" -Duser.home="${JENKINS_HOME}" -Djenkins.install.runSetupWizard="${JENKINS_RUN_SETUP_WIZARD}" -jar "${BIN_DIR}/jenkins.war" --argumentsRealm.roles.user=admin --argumentsRealm.passwd.admin=admin --argumentsRealm.roles.admin=admin >> $${LOG_TO:-/dev/stdout} 2>&1 &

jenkins-restart: ;
