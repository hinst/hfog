import ApiURL from "./ApiURL";

function LoadJson(url, receiver) {
    fetch(ApiURL + url).then(response => response.json().then(data => receiver(data)));
}

export function LoadBugList(receiver) {
    LoadJson("/bugs", receiver);
}

