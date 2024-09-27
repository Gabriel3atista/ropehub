import { header } from '@/lib/links';
import { NextIcon } from '@/components';

export function AppHeader () {
    return (
        header.map((link) => {
            return (
                <p>
                    <NextIcon name="mdi:github" /> - { link.label }
                </p>
            )
        })
    )
}