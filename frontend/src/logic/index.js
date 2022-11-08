import axios from 'axios'


export const fetchMostCommonZipcodes = async () => {
    const res = await axios.get('/common-zipcodes')

    return res.data
}

export const fetchTotalOverlap = async () => {
    const res = await axios.get('/total-overlap')

    return res.data
}

export const fetchOverlapBetweenGroups = async () => {
    const res = await axios.get('/overlap-between-groups')
    return res.data
}

export const fetcAllUsersWithNumberOfElements = async () => {
    const res = await axios.get('/all-users')
    return res.data
}