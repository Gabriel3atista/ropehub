interface Link {
    label: string,
    href?: string,
    icon?: string,
    target?: string,
    children?: Link[]
}

const header: Link[] = [
    {
        label: 'Início',
        href: '/',
        icon: '',
    },
    {
        label: 'Vagas',
        href: '/vagas',
        icon: '',
        target: 'blank',
    },
    {
        label: 'Organizações',
        icon: '',
        children: [
            {
                label: 'Teste',
                href: '/teste',
            }
        ],
    },
]

export {
    header
}