package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	model "github.com/nettica-com/nettica-admin/model"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

///
/// Mongo DB primitives
///

// Serialize write interface to disk
func Serialize(id string, parm string, col string, c interface{}) error {
	//b, err := json.MarshalIndent(c, "", "  ")
	//if err != nil {
	//	return err
	//}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	data, err := json.Marshal(c)
	//	json := fmt.Sprintf("%v", user)
	var b interface{}
	err = bson.UnmarshalExtJSON([]byte(data), true, &b)

	collection := client.Database("nettica").Collection(col)

	findstr := fmt.Sprintf("{\"%s\":\"%s\"}", parm, id)
	var filter interface{}
	err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	update := bson.M{
		"$set": b,
	}

	opts := options.Update().SetUpsert(true)
	_, err = collection.UpdateOne(ctx, filter, update, opts)

	//	if res != nil && res.Err != nil {
	//		collection.InsertOne(ctx, b)
	//	}

	return err
}

// Deserialize read interface from disk
func Deserialize(id string, parm string, col string, t reflect.Type) (interface{}, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection(col)

	findstr := fmt.Sprintf("{\"%s\":\"%s\"}", parm, id)
	var filter interface{}
	err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	switch t.String() {
	case "model.Account":
		var c *model.Account
		err = collection.FindOne(ctx, filter).Decode(&c)
		return c, err

	case "model.VPN":
		var c *model.VPN
		err = collection.FindOne(ctx, filter).Decode(&c)
		return c, err

	case "model.Device":
		var c *model.Device
		err = collection.FindOne(ctx, filter).Decode(&c)
		return c, err

	case "model.User":
		var c *model.User
		err = collection.FindOne(ctx, filter).Decode(&c)
		return c, err

	case "model.Network":
		var c *model.Network
		err = collection.FindOne(ctx, filter).Decode(&c)
		return c, err

	case "model.Subscription":
		var c *model.Subscription
		err = collection.FindOne(ctx, filter).Decode(&c)
		return c, err

	case "model.Service":
		var c *model.Service
		err = collection.FindOne(ctx, filter).Decode(&c)
		return c, err

	case "model.Server":
		var c *model.Server
		err = collection.FindOne(ctx, filter).Decode(&c)
		return c, err
	}
	log.Infof("reflect.TypeOf(t) = %v", t.String())

	return nil, nil
}

// DeleteVPN removes the given id from the given collection
func DeleteVPN(id string, col string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection(col)

	findstr := fmt.Sprintf("{\"id\":\"%s\"}", id)
	var filter interface{}
	err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	collection.FindOneAndDelete(ctx, filter)

	return nil
}

// Delete removes the given id from the given collection
func Delete(id string, ident string, col string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection(col)

	findstr := fmt.Sprintf("{\"%s\":\"%s\"}", ident, id)
	var filter interface{}
	err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	collection.FindOneAndDelete(ctx, filter)

	return nil
}

// ReadAllDevices from MongoDB
func ReadAllDevices(param string, id string) ([]*model.Device, error) {
	devices := make([]*model.Device, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("devices")

	filter := bson.D{}
	if id != "" {
		findstr := fmt.Sprintf("{\"%s\":\"%s\"}", param, id)
		err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	}

	cursor, err := collection.Find(ctx, filter)

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var device *model.Device
			err = cursor.Decode(&device)
			if err == nil {
				devices = append(devices, device)
			}
		}

	}

	return devices, err

}

// ReadAllHosts from MongoDB
func ReadAllVPNs(param string, id string) ([]*model.VPN, error) {
	vpns := make([]*model.VPN, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("vpns")

	filter := bson.D{}
	if id != "" {
		findstr := fmt.Sprintf("{\"%s\":\"%s\"}", param, id)
		err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	}

	cursor, err := collection.Find(ctx, filter)

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var vpn *model.VPN
			err = cursor.Decode(&vpn)
			if err == nil {
				vpns = append(vpns, vpn)
			}
		}

	}

	return vpns, err

}

// ReadAllNetworks from MongoDB
func ReadAllNetworks(param string, id string) []*model.Network {
	nets := make([]*model.Network, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	filter := bson.D{}
	if id != "" {
		findstr := fmt.Sprintf("{\"%s\":\"%s\"}", param, id)
		err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	}

	collection := client.Database("nettica").Collection("network")
	cursor, err := collection.Find(ctx, filter)

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var net *model.Network
			err = cursor.Decode(&net)
			if err == nil {
				nets = append(nets, net)
			}
		}

	}

	return nets

}

// ReadAllUsers from MongoDB
func ReadAllUsers() []*model.User {
	users := make([]*model.User, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("users")

	cursor, err := collection.Find(ctx, bson.D{})

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var user *model.User
			err = cursor.Decode(&user)
			if err == nil {
				users = append(users, user)
			}
		}

	}

	return users

}

// ReadAllAccounts from MongoDB
func ReadAllAccounts(email string) ([]*model.Account, error) {
	accounts := make([]*model.Account, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("accounts")

	filter := bson.D{}
	if email != "" {
		findstr := fmt.Sprintf("{\"%s\":\"%s\"}", "email", email)
		err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	}

	cursor, err := collection.Find(ctx, filter)

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var account *model.Account
			err = cursor.Decode(&account)
			if err == nil {
				accounts = append(accounts, account)
			}
		}

	}

	return accounts, err

}

// ReadAllAccountsForID from MongoDB
func ReadAllAccountsForID(id string) ([]*model.Account, error) {
	accounts := make([]*model.Account, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("accounts")

	filter := bson.D{}
	if id != "" {
		findstr := fmt.Sprintf("{\"%s\":\"%s\"}", "parent", id)
		err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	}

	cursor, err := collection.Find(ctx, filter)

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var account *model.Account
			err = cursor.Decode(&account)
			if err == nil {
				accounts = append(accounts, account)
			}
		}

	}

	return accounts, err

}

// ReadAllSubscriptions from MongoDB
func ReadAllSubscriptions(email string) ([]*model.Subscription, error) {
	subscriptions := make([]*model.Subscription, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("subscriptions")

	filter := bson.D{}
	if email != "" {
		findstr := fmt.Sprintf("{\"%s\":\"%s\"}", "email", email)
		err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	}

	cursor, err := collection.Find(ctx, filter)

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var subscription *model.Subscription
			err = cursor.Decode(&subscription)
			if err == nil {
				subscriptions = append(subscriptions, subscription)
			}
		}

	}

	return subscriptions, err

}

// ReadAllServices from MongoDB
func ReadAllServices(email string) ([]*model.Service, error) {
	services := make([]*model.Service, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("service")

	filter := bson.D{}
	if email != "" {
		findstr := fmt.Sprintf("{\"%s\":\"%s\"}", "email", email)
		err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	}

	cursor, err := collection.Find(ctx, filter)

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var service *model.Service
			err = cursor.Decode(&service)
			if err == nil {
				services = append(services, service)
			}
		}

	}

	return services, err

}

// ReadAllServers from MongoDB
func ReadAllServers() ([]*model.Server, error) {
	servers := make([]*model.Server, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("servers")
	cursor, err := collection.Find(ctx, bson.D{})
	if err == nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var server *model.Server
			err = cursor.Decode(&server)
			if err == nil {
				servers = append(servers, server)
			}
		}
	}
	return servers, err
}

// ReadServiceHost from MongoDB
func ReadServiceHost(id string) ([]*model.Service, error) {
	services := make([]*model.Service, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_CONNECTION_STRING")))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Error(err)
		}
	}()

	collection := client.Database("nettica").Collection("service")

	filter := bson.D{}
	if id != "" {
		findstr := fmt.Sprintf("{\"%s\":\"%s\"}", "serviceGroup", id)
		err = bson.UnmarshalExtJSON([]byte(findstr), true, &filter)

	}

	cursor, err := collection.Find(ctx, filter)

	if err == nil {

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var service *model.Service
			err = cursor.Decode(&service)
			if err == nil {
				services = append(services, service)
			}
		}

	}

	return services, err

}
