from hstest.stage_test import StageTest
from hstest.test_case import TestCase
from hstest.check_result import CheckResult
import re

CheckResult.correct = lambda: CheckResult(True, '')
CheckResult.wrong = lambda feedback: CheckResult(False, feedback)


class LoanCalcTest(StageTest):
    def generate(self):
        return [
            TestCase(
                args=['--principal=1000000', '--periods=60', '--interest=10'],
                attach=21248
            ),
            TestCase(
                args=['--principal=1000000', '--periods=8', '--interest=9.8'],
                attach=129638
            ),
            TestCase(
                args=['--principal=3000000', '--periods=302', '--interest=11.2'],
                attach=29803
            ),
            TestCase(
                args=['--principal=500000', '--payment=23000', '--interest=7.8'],
                attach=[2, 0]
            ),
            TestCase(
                args=['--principal=700000', '--payment=26000', '--interest=9.1'],
                attach=[2, 7]
            ),
            TestCase(
                args=['--payment=8721.8', '--periods=120', '--interest=5.6'],
                attach=(800000,)
            ),
            TestCase(
                args=['--payment=6898.02', '--periods=240', '--interest=3.4'],
                attach=(1200001,)
            )
        ]

    def check(self, reply, attach):
        numbers = re.findall(r'[-+]?(\d*\.\d+|\d+)', reply)
        if len(numbers) == 0:
            return CheckResult.wrong(
                'No numbers in the answer',
            )

        if isinstance(attach, tuple):
            for i in numbers:
                if abs(attach[0] - float(i)) < 2:
                    return CheckResult.correct()
            output = 'Numbers in your answer: ' + ' '.join(numbers)
            output += '. But correct principal is {0}'.format(attach[0])
            return CheckResult.wrong(output)

        if isinstance(attach, list):
            # to exclude answers like 'it takes 2.01 years'
            # but 'it takes 2.0 years' let it be OK.
            epsilon = 0.00001
            numbers = [
                int(float(x)) for x in numbers
                if abs(int(float(x)) - float(x)) < epsilon
            ]
            if attach[1] == 0:
                if 'year' in reply.lower() and attach[0] in numbers:
                    return CheckResult.correct()

                output = 'Correct result: {0} years, but you output "{1}"'
                return CheckResult.wrong(
                    output.format(attach[0], reply),
                )
            else:
                if attach[0] in numbers and 'year' in reply.lower():
                    if attach[1] in numbers and 'month' in reply.lower():
                        return CheckResult.correct()

                output = (
                    'Correct result: {0} years {1} months, '
                    'but you output "{2}"'
                )
                return CheckResult.wrong(
                    output.format(attach[0], attach[1], reply),
                )

        if str(attach) not in reply.lower():
            output = (
                'Correct annuity payment is {0} but you output numbers: {1}'
            )
            figures = ' '.join(numbers)
            return CheckResult.wrong(
                output.format(attach, figures),
            )

        return CheckResult.correct()


if __name__ == '__main__':
    LoanCalcTest('creditcalc.creditcalc').run_tests()
