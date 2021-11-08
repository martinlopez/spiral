import React, {useState} from 'react'
import axios from 'axios'
import _ from 'lodash'
const client = axios.create({
    baseURL: 'http://localhost:8080',
    timeout: 3000,
  })

const Home = () => {
    const [cols, setCols] = useState(1)
    const [rows, setRows] = useState(1)
    const [table, setTable] = useState(undefined)
    const handleChangeCols= (event) => {
        setCols(event.target.value)
    }
    const handleChangeRows= (event) => {
        setRows(event.target.value)
    }
    const handleSubmit= async () => {
        await client.get(`/spiral?cols=${cols}&rows=${rows}`).then(resp  => {
            setTable(resp.data.rows)
        }).catch((e) => {
            console.log(e)
        })     
    }

    return (
        <div style={{"display": "inline-table"}}>
            <h1>Fibonacci Spiral</h1>
            <h4>Matrix properties</h4>
                <label>
                Columns:
                <input type="number" value={cols} onChange={handleChangeCols} />
                </label>
                <label>
                Rows:
                <input type="number" value={rows} onChange={handleChangeRows} />
                </label>
                <input type="submit" value="Submit" onClick={handleSubmit} />
                {table  && <table key={"table"} border="3px"> 
                    <tbody>
                        {table.map(e => <tr key={_.uniqueId()}>
                          { e.map(f => <td key={_.uniqueId()}>{f}</td>)}  
                        </tr>)}     
                    </tbody>
                </table>
                } 
        </div>
    )
}

export default Home