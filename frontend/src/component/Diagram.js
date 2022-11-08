import HighchartsReact from 'highcharts-react-official'
import Highcharts from 'highcharts'
import VennModule from "highcharts/modules/venn.js"
import {useEffect, useState} from 'react'
import {fetcAllUsersWithNumberOfElements, fetchOverlapBetweenGroups, fetchTotalOverlap} from '../logic'
VennModule(Highcharts)

const Diagram = () => {
    let vennOptions = {
        title: {},
        chart: {
            borderWidth: 2,
            borderColor: "#fff",
            borderRadius: 20,
            className : "",
            animation: false,
        },
        series: [
            {
            type: 'venn',
            name: 'Overlap Between Groups',
            data: []
        }]
    }

    const [vennOptionsState, setVennOptionsState] = useState({})

    const fetchData = async () => {
        const data = await fetchOverlapBetweenGroups()
        return data
    }

    const fetchTotalOverlapsData = async () => {
        const data = await fetchTotalOverlap()
        return data
    }

    const fetchUserTypes = async () => {
        const data = await fetcAllUsersWithNumberOfElements()
        return data
    }

    const setDiagramData = async () => {
        const overlapBetweenGroupsData = await fetchData()
        const totalOverlapData = await fetchTotalOverlapsData()
        const allUserTypes = await fetchUserTypes()

        vennOptions.series[0].data = [...vennOptions.series[0].data,
            ...allUserTypes.map(el => ({sets: [el.name], value: el.numberOfElements, name: el.name}))
        ]

        let allGroups = []

        overlapBetweenGroupsData.forEach((elem) => {
            const groups = elem.name.split("->")
            allGroups = [
                ...allGroups,
                {name: groups[0], elementsNumber: elem.firstItemElements.length},
                {name: groups[1], elementsNumber: elem.secondItemElements.length},
            ]
            vennOptions.series[0].data = [...vennOptions.series[0].data, {sets: groups, value: elem.numberOfUsers, name: elem.name}]
        })

        vennOptions.series[0].data = [...vennOptions.series[0].data, {
            sets: [...new Set(allGroups.map(el => el.name))], value: totalOverlapData.totalOverlap.length, name: 'Total Overlap'}
        ]

        vennOptions.series[0].data = [...vennOptions.series[0].data, {
            sets: [...new Set(allGroups.map(el => el.name))], value: totalOverlapData.totalOverlap.length, name: 'Total Overlap'}
        ]

        vennOptions.title.text = `Total Overlap Between Groups is ${totalOverlapData.overlapPercentage}%`

        setVennOptionsState(vennOptions)

    }

    useEffect(() => {
        setDiagramData()
    }, [])

    return (
        <HighchartsReact
            highcharts={Highcharts}
            options={vennOptionsState}
        />
    )
}

export default Diagram