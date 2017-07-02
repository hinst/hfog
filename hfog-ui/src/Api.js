import ApiURL from "./ApiURL";

function LoadJson(url, receiver) {
    fetch(ApiURL + url).then(response => response.json().then(data => receiver(data)));
}

export function LoadBugList(receiver) {
    LoadJson("/bugs?s", receiver);
}

export function LoadBug(bugId, receiver) {
    LoadJson("/getBug?id=" + encodeURIComponent(bugId), receiver);
}

export function LoadBugListFiltered(filterString, receiver) {
    LoadJson("/getBugsFiltered?filter=" + encodeURIComponent(filterString), receiver);
}

