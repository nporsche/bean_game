struct Player{
    1: required int64 id;
    2: required string name;
    3: optional Point location;
}

exception BeanException{
    1: int errNo;
    2: string errMsg;
}

struct Game {
    1: required int64 id;            
    2: required string name;            
    3: required int64 hostPlayerId;
    4: required Rectangle rect;
    5: required int players;
    6: required int maxPlayers;
}

struct Point {
    1: required int64 longitude;
    2: required int64 latitude;
}

struct Rectangle {
    1: required Point pointMin;
    2: required Point pointMax;
}

service GameManagerService {
    //CreateGame returns game id
    int64 CreateGame(1:Player player, 2:int maxPlayers, 3:int cityId, 4:Rectangle rect) throws (1: BeanException excep),

    //List games which belong to this cityId
    List<Game> ListGames(1:int cityId) throws (1: BeanException excep),

    //Usage:
    //1. Join the game.
    //2. If already joined the game, you could query players status.
    //3. Throws the exception if the game quit(errno = 1).
    //4. Throws the exception if the game starts(errno = 2).
    List<Player> JoinGame(1:int64 gameId) throws (1: BeanException excep)
}

service PlayerService {
    //Report my location and returns other players' location
    List<Player> Report(1: int64 game, 2: Player me) throws(1: BeanException excep)
    List<Point> ListBeans(1: int64 game) throws(1: BeanException excep)
} 
