package repositories

import "calcal/pkg/repositories/mongodb"

type Repositories struct {
	MongoDB mongodb.MongoRepository
}
