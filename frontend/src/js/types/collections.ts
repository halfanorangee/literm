interface collections {
    id: string;
    name: string;
    children?: connections[];
}

interface connections {
    ID: number;
    Collection_ID: number;
    ConnName: string;
    IP: string;
    Port: number;
    UserName: string;
    Password: string;
    Key: string;
}