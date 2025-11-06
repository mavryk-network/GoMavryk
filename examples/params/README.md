## Know your block

Use GoMavryk to compute static information about a given block height. Since mainnet has switched from 4096 blocks/cycle to 8192 it has become necessary to have a good baseline set of functions to keep track of such info, in particular, since some protocols contain off-by-1 bugs.

### Usage

```sh
Params Test
  -net string
      simulate with network
  -node string
      node url (default "https://rpc.tzpro.io")
  -proto string
      simulate with protocol
  -v  be verbose
  ```

### Examples

```sh
$ go run . 2244609
Using protocol PtAtLas on Mainnet
Height ......................  2244609
Protocol ....................  PtAtLasdzXg4XxeVNtWheo13nG4wHXP22qYMqFcT3fyBpWkFero
Period ......................  proposal 69
StartCycle ..................  468
StartBlockOffset ............  2244608
VoteBlockOffset .............  0
BlocksPerCycle ..............  8192
BlocksPerVotingPeriod .......  40960
-----------------------------
IsCycleStart ................  true
IsCycleEnd ..................  false
IsSnapshotBlock .............  false
IsSeedRequired ..............  false
CycleFromHeight .............  468
CycleStartHeight ............  2244609
CycleEndHeight ..............  2252800
SnapshotIndex ...............  0
MaxSnapshotIndex ............  15
VotingStartCycleFromHeight ..  468
IsVoteStart .................  true
IsVoteEnd ...................  false
VoteStartHeight .............  2244609
VoteEndHeight ...............  2285568
```

