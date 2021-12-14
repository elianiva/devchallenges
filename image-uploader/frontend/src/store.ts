import { writable } from "svelte/store";

export const isUploading = writable(false);
export const isFinished = writable(false);
