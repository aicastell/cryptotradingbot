Cryptotradingbot

This is the crypto trading bot I am working on to automate buy and sell
operations on different cryto exchanges and with different trading strategies.

Currently I am working to support "bitstamp" exchange with a basic/simple EMA
strategy (Exponential Moving Average).

EMA offers a smooth relation between price and time. I am using this maths:

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

