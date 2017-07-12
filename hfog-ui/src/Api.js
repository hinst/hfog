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

export class FilterArgs {
    constructor() {
        this.filterString = "";
        this.commentsEnabled = false;
    }
}

// Type of args is FilterArgs.
export function LoadBugListFiltered(args, receiver) {
    var url = "/getBugsFiltered?filter=" + encodeURIComponent(args.filterString);
    if (args.commentsEnabled)
        url += "&ce=y";
    LoadJson(url, receiver);
}

