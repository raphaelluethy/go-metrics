"use client";
import * as React from 'react';
import {LineChart, Line, XAxis, YAxis, Tooltip, Legend} from 'recharts';

export default function Chart() {
    const [data, setData] = React.useState([
        {
            average: 400,
            today: 240,
        },
        {
            average: 300,
            today: 139,
        },
        {
            average: 200,
            today: 980,
        },
        {
            average: 278,
            today: 390,
        },
        {
            average: 189,
            today: 480,
        },
        {
            average: 239,
            today: 380,
        },
        {
            average: 349,
            today: 430,
        }]);

    return (
        <div className={"h-screen; w-screen bg-amber-400"}>
            <div className={"bg-green-300"}>
                <LineChart
                    width={500}
                    height={300}
                    data={data}
                    margin={{
                        top: 5,
                        right: 30,
                        left: 20,
                        bottom: 5,
                    }}
                >
                    <XAxis/>
                    <YAxis/>
                    <Tooltip/>
                    <Legend/>
                    <Line type="monotone" dataKey="today" stroke="#8884d8" activeDot={{r: 8}}/>
                    <Line type="monotone" dataKey="average" stroke="#82ca9d"/>
                </LineChart>
            </div>
        </div>
    );
}
