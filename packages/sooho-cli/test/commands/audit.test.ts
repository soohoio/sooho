import {expect, test} from '@oclif/test'

describe('audit', () => {
  test
  .stdout()
  .stderr()
  .command(['audit', 'test/commands/Vulnerable.sol'])
  .it('parse file', ctx => {
    expect(ctx.stderr.includes('Parse files')).to.equal(true)
  })

  test
  .stdout()
  .stderr()
  .command(['audit', 'test/commands/Vulnerable.sol'])
  .it('detect vulnerabilities', ctx => {
    expect(ctx.stderr.includes('Vulnerabilities have detected!')).to.equal(true)
    expect(ctx.stdout).to.equal([
      '┌────────────────────┬──────────┬──────────┬──────────────────────────────────────────────────┐',
      '│ CVE ID             │ Type     │ Severity │ Desc                                             │',
      '├────────────────────┼──────────┼──────────┼──────────────────────────────────────────────────┤',
      '│ CVE-2018-12056     │ SWC-120  │ 7.5      │ The maxRandom function of a smart contract impl… │',
      '│                    │          │          │                                                  │',
      '└────────────────────┴──────────┴──────────┴──────────────────────────────────────────────────┘',
      ''
    ].join('\n'))
  })
})
