# ITF-Diff Markdown

## Step 1 to Step 2

<details open>

<summary>Variables</summary>

<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.coins`|`None`|`<<[ amount \|-> 2, denom \|-> "atom" ], [ amount \|-> 1, denom \|-> "muon" ]>>`|
|`action.receiver`|`None`|`"Bob"`|
|`action.sender`|`None`|`"Bob"`|
|`action.tag`|`init`|`send`|
|`action.balances`|`SetAsFun({<<"Alice", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Bob", SetAsFun({<<"atom", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"gluon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>, <<"muon", 57896044618658097711785492504343953926634992332820282019728792003956564819967>>})>>, <<"Carol", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Dave", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>, <<"Eve", SetAsFun({<<"atom", 0>>, <<"gluon", 0>>, <<"muon", 0>>})>>})`|`None`|

</details>
<details open>

<summary><code>outcome</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`outcome`|``|`SUCCESS`|

</details>
<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`0`|`1`|

</details>

</details>

## Step 2 to Step 3

<details open>

<summary>Variables</summary>

<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`1`|`2`|

</details>
<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.coins[0]`|`[ amount \|-> 2, denom \|-> "atom" ]`|`[ amount \|-> 3, denom \|-> "atom" ]`|
|`action.coins[1]`|`[ amount \|-> 1, denom \|-> "muon" ]`|`[ amount \|-> 2, denom \|-> "muon" ]`|

</details>

</details>

## Step 3 to Step 4

<details open>

<summary>Variables</summary>

<details open>

<summary><code>step</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`step`|`2`|`3`|

</details>
<details open>

<summary><code>action</code></summary>


|KeyPath|Old|New|
|-|-|-|
|`action.coins[0]`|`[ amount \|-> 3, denom \|-> "atom" ]`|`[ amount \|-> 57896044618658097711785492504343953926634992332820282019728792003956564819965, denom \|-> "atom" ]`|
|`action.coins[1]`|`[ amount \|-> 2, denom \|-> "muon" ]`|`[ amount \|-> 300, denom \|-> "muon" ]`|

</details>

</details>
