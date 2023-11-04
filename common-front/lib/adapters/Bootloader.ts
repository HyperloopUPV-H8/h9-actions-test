export type BootloaderUploadRequest = { board: string; file: string };
export type BootloaderUploadResponse = { percentage: number; failure: boolean };

export type BootloaderDownloadRequest = { board: string };
export type BootloaderDownloadResponse = {
    percentage: number;
    failure: boolean;
    file: string;
};
