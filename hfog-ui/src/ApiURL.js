const devMode = !process.env.NODE_ENV || process.env.NODE_ENV === 'development';

const ApiURL = devMode ? "http://localhost:9000/FogBugzBackupApi" : "/FogBugzBackupApi";

export default ApiURL;