"""Create and test simple job scheduler to run a function after some milliseconds pass."""
import threading
import time
from collections.abc import Callable

type Job = tuple[Callable[[], None], int]

ms_in_ns: int = 1000 * 1000
ns_in_s: int = 1000 * 1000 * 1000


class JobScheduler:
    """Simple job scheduler."""

    __jobs: list[Job]

    def __init__(self):
        """Init simple job scheduler that has main thread to schedule and one more thread to run jobs."""
        self.__jobs = []
        self.__cv = threading.Condition()
        threading.Thread(target=self.__poll).start()

    def __poll(self):
        while True:
            to_run: list[Callable[[], None]] = []
            not_yet: list[Job] = []

            with self.__cv:
                now_ns = time.time_ns()
                for job in self.__jobs:
                    if job[1] <= now_ns:
                        to_run.append(job[0])
                    else:
                        not_yet.append(job)
                self.__jobs = not_yet

            if to_run:
                for func in to_run:
                    func()

            with self.__cv:
                if not self.__jobs:
                    self.__cv.wait_for(lambda: self.__jobs)
                else:
                    ns_remaining = min(due for fn, due in self.__jobs) - time.time_ns()
                    if ns_remaining > 0:
                        self.__cv.wait(ns_remaining / ns_in_s)

    def schedule(self, job: Callable[[], None], after_ms: int):
        """Schedule a function to run after after_ms."""
        with self.__cv:
            self.__jobs.append((job, time.time_ns() + after_ms * ms_in_ns))
            self.__cv.notify_all()


def main():
    """Test job scheduler."""
    test_sch = JobScheduler()
    test_sch.schedule(lambda: print("First after 1s."), 1000)
    test_sch.schedule(lambda: print("Second after 2s."), 2000)
    test_sch.schedule(lambda: print("First after 0.5s."), 500)
    test_sch.schedule(lambda: print("First after 0.25s."), 250)
    test_sch.schedule(lambda: print("First after 0s."), 0)


if __name__ == "__main__":
    main()
