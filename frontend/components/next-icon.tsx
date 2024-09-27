'use client'

import { Icon } from '@iconify-icon/react';

interface Props {
  name: string
}

export function NextIcon (props: Props) {
  return <Icon icon={ props.name } />
}