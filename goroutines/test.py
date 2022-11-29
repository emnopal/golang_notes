# Sama aja kaya python ternyata, makan ram 2 gb juga hahaha

import asyncio


async def DisplayNum(num: int) -> None:
    print("Display", num)
    await asyncio.sleep(10)


async def TestHowLightIs() -> None:
    await asyncio.gather(*(DisplayNum(i) for i in range(1000000)))

if __name__ == '__main__':
    asyncio.run(TestHowLightIs())
