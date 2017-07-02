export function GetURL() {
    var urlParams = new URLSearchParams(window.location.search);
    return "AccessKey=" + encodeURIComponent(urlParams.get("AccessKey"));
}