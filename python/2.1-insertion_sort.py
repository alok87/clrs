import sys
import random
import logging
import time
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


def swap(a, b):
    temp = b
    b = a
    a = temp
    return a, b


if __name__ == '__main__':
    n = 1000
    arr = random.sample(xrange(n), 100)
    logger.info('Insertion sort %s' % arr)

    j = 1
    while j <= len(arr)-1:
        i = 0
        while i < j:
            if arr[i] > arr[j]:
                arr[i], arr[j] = swap(arr[i], arr[j])
            i += 1
        j += 1
        logger.debug(arr)

    logger.info('Sorted list %s' % arr)
