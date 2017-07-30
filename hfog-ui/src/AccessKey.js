export function GetURL() {
    const urlParams = new URLSearchParams(window.location.search);
    let accessKey = urlParams.get("AccessKey");
    if (accessKey == null)
        accessKey = "";
    return "AccessKey=" + encodeURIComponent(accessKey);
}