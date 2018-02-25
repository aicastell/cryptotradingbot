
Esto es un bot para comprar y vender automatico utilizando el algoritmo EMA

Media móvil exponencial (Exponential Moving Average)

La media móvil exponencial, al igual que la media móvil simple, ofrece una
correlación suavizada entre la acción del precio y el transcurso del tiempo. La
diferencia está en que el cálculo de la media móvil exponencial da más
importancia a los últimos datos obtenidos durante un determinado período.

La fórmula:

    EMA(t) = EMA(t – 1) + K*[Precio(t) – EMA(t – 1)]

Donde:

    t = valor actual

    t-1 = valor previo

    n = window_len

    K = 2 / (n + 1)


