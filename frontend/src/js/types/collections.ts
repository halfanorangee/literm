interface collections {
    id: string;
    name: string;
    children?: connections[];
}

interface connections {
    id: number;
    conn_name: string;
    ip: string;
    port: number;
    user_name: string;
    password: string;
    key: string;
}

export async function getConnInfos(): Promise<ConnInfo[]> {
    try {
        const connInfos: ConnInfo[] = await window.go.App.GetConnInfos();
        console.log('Fetched ConnInfos:', connInfos);
        return connInfos;
    } catch (error) {
        console.error('Error fetching conn infos:', error);
        return [];
    }
}