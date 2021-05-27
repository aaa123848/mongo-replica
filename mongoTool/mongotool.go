package mongotool

import (
	"context"
	"errors"
	
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context = context.Background()

type MongoTool struct {
	Client *mongo.Client
}


func (m *MongoTool) Connect(url string){
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil{
		panic(err)
	}
	m.Client = client
}

func (m *MongoTool) Disconnect(){
	if err := m.Client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func (m *MongoTool) UpdateWithSession(db string, col string, filter bson.D, update bson.D) error{
	sess, err := m.Client.StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(context.TODO())
	err = mongo.WithSession(context.TODO(), sess, func(sessCtx mongo.SessionContext)error{
		if err := sess.StartTransaction(); err != nil{
			return err
		}
		coll := m.Client.Database(db).Collection(col)

		res, err := coll.UpdateOne(
			sessCtx,
			filter,
			update,
		)
		if err != nil{
			_ = sess.AbortTransaction(context.Background())
			return err
		}
		if res.ModifiedCount < 1{
			_ = sess.AbortTransaction(context.Background())
			return errors.New("ALready exist")
		}
		return sess.CommitTransaction(context.Background())
	})
	if err != nil{
		return err
	}
	return nil
}

