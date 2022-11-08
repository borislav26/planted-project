import React, {useEffect, useState} from 'react'
import {fetchMostCommonZipcodes} from '../logic'

import './style.css'

const CommonZipcodes = () => {
    const [zipcodes, setZipcodes] = useState([])
    const fetchData = async () => {
        const data = await fetchMostCommonZipcodes()
        let counter = 0
        const finalData = data.map(el => ({...el, rank: ++counter}))
        setZipcodes(finalData)
    }

    useEffect(   () => {
        fetchData()
    }, [])

    return (
        <>
            <div className='CommonZipcodesTitle'>
                <h1>Zipcodes Ranking</h1>
            </div>
            <div className="CommonZipcodes">
                <table>
                    <tr>
                        <th>Zipcode</th>
                        <th>Total Number</th>
                        <th>Rank</th>
                    </tr>
                    {zipcodes.map(el=> {
                        return (
                            <tr key={el.key}>
                                <td>{el.key}</td>
                                <td>{el.value}</td>
                                <td>{el.rank}</td>
                            </tr>
                        )
                    })}
                </table>
            </div>
        </>
    )
}

export default CommonZipcodes