<div align="center"><img src="./docs/images/charonlogo.svg" /></div>
<h1 align="center">Charon<br/>Промежуточный клиент Distributed Validator</h1>

<p align="center"><a href="https://github.com/obolnetwork/charon/releases/"><img src="https://img.shields.io/github/tag/obolnetwork/charon.svg"></a>
<a href="https://github.com/ObolNetwork/charon/blob/main/LICENSE"><img src="https://img.shields.io/github/license/obolnetwork/charon.svg"></a>
<a href="https://godoc.org/github.com/obolnetwork/charon"><img src="https://godoc.org/github.com/obolnetwork/charon?status.svg"></a>
<a href="https://goreportcard.com/report/github.com/obolnetwork/charon"><img src="https://goreportcard.com/badge/github.com/obolnetwork/charon"></a>
<a href="https://github.com/ObolNetwork/charon/actions/workflows/golangci-lint.yml"><img src="https://github.com/obolnetwork/charon/workflows/golangci-lint/badge.svg"></a></p>

Это репозиторий содержит исходный код клиента распределённого валидатора _Charon_ (произносится как "харон"); HTTP промежуточный клиент для Ethereum Staking, который позволяет безопасно запускать один валидатор через группу независимых узлов (нод).

Charon дополняется веб-приложением [Distributed Validator Launchpad](https://goerli.launchpad.obol.tech/) для создания ключей распределённого валидатора.

Charon используется стейкерами для разделения функций Ethereum Validators между различными инстансами и клиентскими версиями.

![Пример кластера Obol](./docs/images/DVCluster.png)

###### Распределенный кластер валидаторов, использующий клиент Charon для снижения риска сбоев клиентов и оборудования

## Краткое описание

Самый простой способ протестировать Charon - использовать репозиторий [charon-distributed-validator-cluster](https://github.com/ObolNetwork/charon-distributed-validator-cluster), который содержит установщик Docker Compose для запуска полного кластера Charon на вашей локальной машине.

## Документация

Веб-сайт [Obol Docs](https://docs.obol.tech/) - лучшее место для начала работы. Важными разделами являются [Вступление](https://docs.obol.tech/docs/intro), [Основные понятия](https://docs.obol.tech/docs/int/key-concepts) и [Charon](https://docs.obol.tech/docs/dv/introducing-charon).

Подробную документацию по этому репозиторию смотрите в папке [docs](docs):

- [Конфигурация](docs/configuration.md): Настройка узла (ноды) charon
- [Архитектура](docs/architecture.md): Обзор архитектуры кластера и узлов (нод) charon
- [Структура проекта](docs/structure.md): Структура папок проекта
- [Ветвление и версии релиза](docs/branching.md): Модель ветвления и модель версий релиза Git
- [Помощь по Go](docs/goguidelines.md): Помощь и рекомендации по разработке на Go
- [Вклад](docs/contributing.md): Как вносить вклад в Charon; githooks, шаблоны PR и т.д.

Есть [charon godocs](https://pkg.go.dev/github.com/obolnetwork/charon) с документацией по исходному коду.

## Поддерживаемые клиенты

Charon интегрируется в стек консенсуса Ethereum в качестве промежуточного программного обеспечения между клиентом валидатора (Validator Client) и нодой маяка (Beacon Node) через официальный [Eth Beacon Node REST API] (https://ethereum.github.io/beacon-APIs/#/). Charon поддерживает любую вышестоящую ноду маяка, которая работает с Beacon API. Charon нацелен на поддержку любого нижестоящего автономного клиента валидатора, который использует Beacon API.

| Клиент                                             |  Beacon Node	| Validator Client | Заметки                                 |
| -------------------------------------------------- | :---------: |  :--------------: |-----------------------------------------|
| [Teku](https://github.com/ConsenSys/teku)          |     ✅      |        ✅        | Полностью поддерживается                |
| [Lighthouse](https://github.com/sigp/lighthouse)   |     ✅      |        ✅        | Полностью поддерживается                |
| [Lodestar](https://github.com/ChainSafe/lodestar)  |     ✅      |       \*️⃣        | Проблема совместимости с DVT            |
| [Vouch](https://github.com/attestantio/vouch)      |     \*️⃣     |        ✅        | Только клиент валидатора                |
| [Prysm](https://github.com/prysmaticlabs/prysm)    |     ✅      |        🛑        | Клиент валидатора требует API gRPC      |
| [Nimbus](https://github.com/status-im/nimbus-eth2) |     ✅      |        ✅        | Скоро будет поддерживаться              |

## Статус проекта

Сеть Obol находится в самом начале своего развития и активно развивается.
Мы движемся быстро, поэтому регулярно заглядывайте сюда, чтобы следить за изменениями.

Charon - это распределенный валидатор, поэтому его основная обязанность - выполнение функций валидации.
В следующей таблице указано, какие функции в публичном тестнете выполняют те или иные клиенты, и какие ещё находятся в стадии разработки (🚧 )

| Функция \ Клиент                        |                      Teku                      |                    Lighthouse                    | Lodestar | Nimbus | Vouch | Prysm |
|--------------------------------------|:----------------------------------------------:|:------------------------------------------------:|:--------:|:------:|:-----:|:-----:|
| _Attestation_                        |                       ✅                        |                        ✅                         |    🚧    |   🚧   |  ✅   |  🚧   |
| _Attestation Aggregation_            |                       🚧                       |                        🚧                        |    🚧    |   🚧   |  🚧   |  🚧   |
| _Block Proposal_                     |                       ✅                        |                        ✅                         |    🚧    |   🚧   |  🚧   |  🚧   |
| _Blinded Block Proposal (mev-boost)_ | [✅](https://ropsten.beaconcha.in/block/555067) | [✅](https://ropsten.etherscan.io/block/12822070) |    🚧    |   🚧   |  🚧   |  🚧   |
| _Sync Committee Message_             |                       ✅                        |                        ✅                         |    🚧    |   🚧   |  🚧   |  🚧   |
| _Sync Committee Contribution_        |                       🚧                       |                        🚧                        |    🚧    |   🚧   |  🚧   |  🚧   |
