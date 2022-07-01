from dataclasses import dataclass, field
from typing import Optional
import matplotlib.pyplot as plt
import pandas as pd
import matplotlib.dates as mdates
import numpy as np
import datetime


@dataclass
class Runner:    
    fatigue: float = field(init=False)
    form: float = field(init=False)
    fitness: float = field(init=False)
    tss: list = field(default_factory=list)
    dates: list = field(default_factory=list)

    def __post_init__(self):

        today = datetime.datetime.today().date()

        date_map_fit = {(today - datetime.timedelta(i)).strftime("%Y-%m-%d"): 0 for i in range(42)}
        date_map_fatigue = {(today - datetime.timedelta(i)).strftime("%Y-%m-%d"): 0 for i in range(7)}
        self.fitness = 0
        self.fatigue = 0
        
        for idx, date in enumerate(self.dates):
            if date in date_map_fit:
                self.fitness += self.tss[idx]
            if date in date_map_fatigue:
                self.fatigue += self.tss[idx]
                
        self.fitness = self.fitness/42
        self.fatigue = self.fatigue/7
        self.form = self.fitness - self.fatigue
    
    def plotTSS(self):
        datesList = [pd.to_datetime(day) for day in self.dates]
        plt.scatter(datesList, self.tss, c='red')
        
        plt.show()
