#!/usr/bin/python3
import csv
import argparse
from stats import Runner


def main(file):
    reader = csv.DictReader(file, delimiter=',')

    tss = []
    dates = []

    for row in reader:
        tss.append(float(row['TSS']))
        dates.append(row["WorkoutDay"])

    obj = Runner(tss, dates)

    print(f"Fitness\n>\t{obj.fitness}\nFatigue\n>\t{obj.fatigue}\nForm\n>\t{obj.form}")

    obj.plotTSS()

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('file', type=argparse.FileType('r'))
    args = parser.parse_args()
    main(args.file)
