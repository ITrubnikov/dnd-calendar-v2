
import ReactDOM from 'react-dom/client';
import './index.css';
import React, { useEffect, useState } from 'react';

// class Square extends React.Component {
//
//     constructor(props) {
//         super(props);
//         this.state ={
//             value: null,
//         };
//     }
//     render() {
//         return (
//             <button
//                 className="square"
//                 onClick={()=> this.props.onClick()}
//                 >
//                 {this.props.value}
//             </button>
//         );
//     }
// }
function App() {
    const [data, setData] = useState([]);

    useEffect(() => {
        fetch('http://localhost:8080/Character') // Укажите URL вашего API
            .then(response => response.json())
            .then(json => setData(json.data))
            .catch(error => console.error('Ошибка при получении данных:', error));
    }, []); // Пустой массив зависимостей означает, что эффект будет выполняться один раз при монтировании компонента

    if (!data) return <div>Загрузка...</div>;

    return (


        <div>

            <h1>Полученные данные</h1>
            {/*<pre>{JSON.stringify(data, null, 2)}</pre>*/}
            <table className="table">
                <thead>
                <tr>
                    <th>ID</th>
                    <th>Name</th>
                    <th>Class</th>
                    <th>Level</th>
                    <th>Mony</th>

                    {/* Добавьте дополнительные заголовки столбцов здесь */}
                </tr>
                </thead>
                <tbody>
                {data.map(item => (
                    <tr key={item.id}>
                        <td>{item.name}</td>
                        <td>{item.class}</td>
                        <td>{item.level}</td>
                        <td>{item.mony}</td>
                        {/* Выведите дополнительные данные здесь */}
                    </tr>
                ))}
                </tbody>
            </table>
        </div>

    )
}

export default App;

function Square(props) {
    return (
        <button className="square" onClick={props.onClick}>
            {props.value}
        </button>
    )
}


class Board extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            squares: Array(9).fill(null),
        };
    }

    handleClick(i) {
        const squares = this.state.squares.slice();
        squares[i] = 'X';
        this.setState({squares:squares})
    }
    renderSquare(i) {
        return( <Square value={this.state.squares[i]}
        onClick={()=> this.handleClick(i)}
        />
        );

    }

    render() {
        const status = 'Next player: X';

        return (
            <div>
                <div className="status">{status}</div>
                <div className="board-row">
                    {this.renderSquare(0)}
                    {this.renderSquare(1)}
                    {this.renderSquare(2)}
                </div>
                <div className="board-row">
                    {this.renderSquare(3)}
                    {this.renderSquare(4)}
                    {this.renderSquare(5)}
                </div>
                <div className="board-row">
                    {this.renderSquare(6)}
                    {this.renderSquare(7)}
                    {this.renderSquare(8)}
                </div>
                <div className="board-row">
                    {this.renderSquare(9)}
                    {this.renderSquare(10)}
                    {this.renderSquare(11)}
                </div>
            </div>
        );
    }
}

class Game extends React.Component {
    render() {
        return (
            <div className="game">
                <div className="game-board">
                    <Board/>
                </div>
                <div className="game-info">
                    <div>{/* status */}</div>
                    <ol>{/* TODO */}</ol>
                </div>
                {/*<div className="my-custome-class">*/}
                {/*    <Square value={"моя самая лучшая кнопка"}/>*/}
                {/*</div>*/}

            </div>

        );
    }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<Game />);
root.render(<App />)
