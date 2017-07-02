import ApiURL from "./ApiURL";
import * as AccessKey from './AccessKey';

function LoadJson(url, receiver) {
    var fullURL = ApiURL + url + "&" + AccessKey.GetURL();
    fetch(fullURL).then(response => response.json().then(data => receiver(data)));
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

