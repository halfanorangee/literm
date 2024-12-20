export interface Menus {
    id: string;
    pid?: string;
    icon?: string;
    index: string;
    title: string;
    permiss?: string;
    children?: Menus[];
}

export const menuInfo: Menus[] = [
    {
        id: '0',
        title: '连接集合',
        index: '/collections',
        icon: 'Management',
    },
    {
        id: '1',
        title: '终端',
        index: '/terminals',
        icon: 'Platform',
    },
    {
        id: '2',
        title: '设置',
        index: '/settings',
        icon: 'Tools',
    }
]