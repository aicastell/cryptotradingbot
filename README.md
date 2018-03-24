# Cryptotradingbot

This is the crypto trading bot I am working on to automate buy and sell
operations on different exchanges.


## Trading Strategies

Cryptotradingbot is being designed to accept different trading strategies.
Currently it supports:

    * strategy01: basic EMA (Exponential Moving Average)
    * strategy02: advanced EMA + RSI (Relative Strength Index)


## Bot configuration

Cryptotradingbot configuration can be set with a JSON file in this format:

    $ cat config.json
    {
        "global": {
            "strategy": "strategy02",
            "fetcher": "poloniex",
            "buycoin": "btc",
            "sellcoin": "eur",
            "invest": 1000,
            "fee": 0.25,
            "period": 60,
            "training_iters": 90
        },
        "ema": {
            "win_len_min": 11,
            "win_len_max": 24
        },
        "rsi": {
            "win_len": 14,
            "buy_level": 40.0,
            "sell_level": 60.0
        }
    }


## Finantial Indicators

Currently two finantial indicators are implemented: EMA and RSI.


### Exponential Moving Average (EMA)

The Exponential Moving Average offers a smooth relation between price and time. I am using this maths:

    EMA(t) = EMA(t – 1) + K*[Precio(t) – EMA(t – 1)]

Where:

    t = current time
    t-1 = previous time
    n = window_len
    K = 2 / (n + 1)

The strategy01 uses two EMA functions (n=13 and n=34) with an interval of 15
minutes between t and (t-1):

    * EMA_13 > EMA_34 ==> climbing price
    * EMA_13 < EMA_34 ==> falling price

The strategy01 consist of:

    * Buy when  (EMA_13(t) < EMA_34(t)) AND (EMA13(t+1) > EMA34(t+1))
    * Sell when (EMA_13(t) > EMA_34(t)) AND (EMA13(t+1) < EMA34(t+1))

You should play with tradingview.com to see the behaviour of this setup.


### Relative Strength Index (RSI)

The Relative Strength Index (RSI) is a momentum oscillator that measures the
speed and change of price movements. The RSI oscillates between zero and 100.
Traditionally the RSI is considered overbought when above 70 and oversold when
below 30.



## Compilation

You need a Go compiler:

    $ go build
    $ go install

After you can start bot with:

    $ bot -c /path/to/config.json

