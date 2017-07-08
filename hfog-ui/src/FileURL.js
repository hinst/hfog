const devMode = !process.env.NODE_ENV || process.env.NODE_ENV === 'development';

const FileURL = devMode ? "" : "/FogBugzBackup";

export default FileURL;