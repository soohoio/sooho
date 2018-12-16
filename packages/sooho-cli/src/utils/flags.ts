import {flags} from '@oclif/command'

export const abstract = flags.boolean({
  char: 'a',
  description: 'turn on abstraction mode',
  default: false
})

export const help = flags.boolean({
  char: 'h',
  description: 'show CLI help'
})

export const save = flags.boolean({
  char: 's',
  description: 'save encrypted result',
  default: false
})
