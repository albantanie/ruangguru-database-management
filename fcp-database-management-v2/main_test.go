package main_test

import (
	"fmt"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"a21hc3NpZ25tZW50/service"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Education Management", func() {
	var studentRepo repo.StudentRepository
	var userRepo repo.UserRepository
	var sessionRepo repo.SessionsRepository
	var classRepo repo.ClassRepository

	var sessionService service.SessionService

	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
		Schema:       "public",
	}
	conn, err := db.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	studentRepo = repo.NewStudentRepo(conn)
	userRepo = repo.NewUserRepo(conn)
	sessionRepo = repo.NewSessionRepo(conn)
	classRepo = repo.NewClassRepo(conn)

	sessionService = service.NewSessionService(sessionRepo)

	BeforeEach(func() {
		err = conn.Migrator().DropTable("students", "users", "sessions", "classes")
		Expect(err).ShouldNot(HaveOccurred())

		conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Student{}, &model.Class{})

		err = db.Reset(conn, "students")
		err = db.Reset(conn, "users")
		err = db.Reset(conn, "sessions")
		err = db.Reset(conn, "classes")
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Repository", func() {

		Describe("Repository", func() {
			Describe("Users repository", func() {
				When("add new user to users table in database postgres", func() {
					It("should save data user to users table indatabase postgres", func() {
						user := model.User{
							Username: "aditira",
							Password: "!opensesame",
						}
						err := userRepo.Add(user)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.User{}
						conn.Model(&model.User{}).First(&result)
						Expect(result.Username).To(Equal(user.Username))
						Expect(result.Password).To(Equal(user.Password))

						err = db.Reset(conn, "users")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})

				When("check user availability in users table database postgres", func() {
					It("return error if present and nil if not present", func() {
						user := model.User{}
						err := userRepo.CheckAvail(user)
						Expect(err).Should(HaveOccurred())

						user = model.User{
							Username: "aditira",
							Password: "!opensesame",
						}

						err = userRepo.Add(user)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.User{}
						conn.Model(&model.User{}).First(&result)
						Expect(result.Username).To(Equal(user.Username))
						Expect(result.Password).To(Equal(user.Password))

						err = userRepo.CheckAvail(user)
						Expect(err).ShouldNot(HaveOccurred())

						err = db.Reset(conn, "users")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})
			})

			Describe("Sessions repository", func() {
				When("add session data to sessions table database postgres", func() {
					It("should save data session to sessions table database postgres", func() {
						session := model.Session{
							Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
							Username: "aditira",
							Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
						}
						err := sessionRepo.AddSessions(session)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						err = db.Reset(conn, "sessions")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})

				When("delete selected session to sessions table database postgres", func() {
					It("should delete data session target from sessions table database postgres", func() {
						session := model.Session{
							Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
							Username: "aditira",
							Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
						}
						err := sessionRepo.AddSessions(session)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						err = sessionRepo.DeleteSession("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).ShouldNot(HaveOccurred())

						result = model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(result).To(Equal(model.Session{}))

						err = db.Reset(conn, "sessions")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})

				When("update selected session to sessions table database postgres", func() {
					It("should update data session target the username field from sessions table database postgres", func() {
						session := model.Session{
							Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
							Username: "aditira",
							Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
						}
						err := sessionRepo.AddSessions(session)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						sessionUpdate := model.Session{
							Token:    "cc03dbac-4085-22ba-75fe-103f9a01b6d5",
							Username: "aditira",
							Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
						}
						err = sessionRepo.UpdateSessions(sessionUpdate)
						Expect(err).ShouldNot(HaveOccurred())

						result = model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(result.Token).To(Equal(sessionUpdate.Token))

						err = db.Reset(conn, "sessions")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})

				When("check expired session with exprired session is 5 hours from now", func() {
					It("should return a session model with token, name, and expired time", func() {
						session := model.Session{
							Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
							Username: "aditira",
							Expiry:   time.Now().Add(5 * time.Hour),
						}
						err := sessionRepo.AddSessions(session)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						tokenFound, err := sessionService.TokenValidity("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).ShouldNot(HaveOccurred())
						Expect(tokenFound.Token).To(Equal("cc03dbea-4085-47ba-86fe-020f5d01a9d8"))
						Expect(tokenFound.Username).To(Equal("aditira"))

						err = db.Reset(conn, "sessions")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})

				When("check expired session with exprired session is 5 hours ago", func() {
					It("should return error message token is expired and empty session model", func() {
						session := model.Session{
							Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
							Username: "aditira",
							Expiry:   time.Now().Add(-25 * time.Hour),
						}
						err := sessionRepo.AddSessions(session)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						tokenFound, err := sessionService.TokenValidity("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).To(Equal(fmt.Errorf("Token is Expired!")))
						Expect(tokenFound).To(Equal(model.Session{}))

						err = db.Reset(conn, "sessions")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})

				When("check session availability with name", func() {
					It("return data session with target name", func() {
						err := sessionRepo.SessionAvailName("aditira")
						Expect(err).Should(HaveOccurred())

						session := model.Session{
							Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
							Username: "aditira",
							Expiry:   time.Now().Add(5 * time.Hour),
						}
						err = sessionRepo.AddSessions(session)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						err = sessionRepo.SessionAvailName("aditira")
						Expect(err).ShouldNot(HaveOccurred())

						err = db.Reset(conn, "sessions")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})

				When("check session availability with token", func() {
					It("return data session with target token", func() {
						_, err := sessionRepo.SessionAvailToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).Should(HaveOccurred())

						session := model.Session{
							Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
							Username: "aditira",
							Expiry:   time.Now().Add(5 * time.Hour),
						}
						err = sessionRepo.AddSessions(session)
						Expect(err).ShouldNot(HaveOccurred())

						result := model.Session{}
						conn.Model(&model.Session{}).First(&result)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						res, err := sessionRepo.SessionAvailToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).ShouldNot(HaveOccurred())
						Expect(res.Token).To(Equal(session.Token))
						Expect(res.Username).To(Equal(session.Username))

						err = db.Reset(conn, "sessions")
						Expect(err).ShouldNot(HaveOccurred())
					})
				})

			})

		})

		Describe("Student repository", func() {
			When("add student data to students table database postgres", func() {
				It("should save student data to students table database postgres", func() {
					student := model.Student{
						Name:    "John",
						Address: "Jl. Raya",
						ClassId: 1,
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Student{}
					conn.Model(&model.Student{}).First(&result)

					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.ClassId).To(Equal(student.ClassId))

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("read all student data from students table database postgres", func() {
				It("should return a list of student data", func() {
					student1 := model.Student{
						Name:    "John",
						Address: "Jl. Raya",
						ClassId: 1,
					}
					err := studentRepo.Store(&student1)
					Expect(err).ShouldNot(HaveOccurred())

					student2 := model.Student{
						Name:    "Doe",
						Address: "Jl. Melati",
						ClassId: 2,
					}
					err = studentRepo.Store(&student2)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchAll()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result).To(HaveLen(2))
					Expect(result[0].Name).To(Equal(student1.Name))
					Expect(result[1].Name).To(Equal(student2.Name))

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("adding new student data to students table in the database", func() {
				It("should save the new student data to students table in the database", func() {
					student := model.Student{
						Name:    "John",
						Address: "123 Main St",
						ClassId: 1,
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Student{}
					conn.Model(&model.Student{}).First(&result)
					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.ClassId).To(Equal(student.ClassId))

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("fetching all student data from students table in the database", func() {
				It("should return a list of all student data", func() {
					students := []model.Student{
						{Name: "John", Address: "123 Main St", ClassId: 1},
						{Name: "Jane", Address: "456 Park Ave", ClassId: 2},
						{Name: "James", Address: "789 Broadway", ClassId: 3},
					}

					for _, student := range students {
						err := studentRepo.Store(&student)
						Expect(err).ShouldNot(HaveOccurred())
					}

					result, err := studentRepo.FetchAll()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(result)).To(Equal(len(students)))

					for i, student := range students {
						Expect(result[i].Name).To(Equal(student.Name))
						Expect(result[i].Address).To(Equal(student.Address))
						Expect(result[i].ClassId).To(Equal(student.ClassId))
					}

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("fetching a single student data by id from students table in the database", func() {
				It("should return a single student data", func() {
					student := model.Student{Name: "John", Address: "123 Main St", ClassId: 1}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.ClassId).To(Equal(student.ClassId))

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("updating student data in students table in the database", func() {
				It("should update the existing student data in students table in the database", func() {
					student := model.Student{Name: "John", Address: "123 Main St", ClassId: 1}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					newStudent := model.Student{Name: "Jane", Address: "456 Park Ave", ClassId: 2}
					err = studentRepo.Update(1, &newStudent)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Student{}
					conn.Model(&model.Student{}).First(&result)
					Expect(result.Name).To(Equal(newStudent.Name))
					Expect(result.Address).To(Equal(newStudent.Address))
					Expect(result.ClassId).To(Equal(newStudent.ClassId))

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("deleting student data in students table in the database", func() {
				It("should delete the existing student data in students table in the database", func() {
					student := model.Student{Name: "John", Address: "123 Main St", ClassId: 1}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					err = studentRepo.Delete(1)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Student{}
					conn.Model(&model.Student{}).First(&result)
					Expect(result).To(Equal(model.Student{}))

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("there are students with classes in the DB", func() {
				BeforeEach(func() {
					class := model.Class{
						Name:       "Mathematics",
						Professor:  "Dr. Smith",
						RoomNumber: 101,
					}

					err := conn.Create(&class).Error
					Expect(err).ShouldNot(HaveOccurred())
				})

				It("should return a list of students with their associated class information", func() {
					expected := []model.StudentClass{
						{
							Name:       "Jane Doe",
							Address:    "123 Main St",
							ClassName:  "Mathematics",
							Professor:  "Dr. Smith",
							RoomNumber: 101,
						},
					}

					student := model.Student{Name: "Jane Doe", Address: "123 Main St", ClassId: 1}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					actual, err := studentRepo.FetchWithClass()
					Expect(err).NotTo(HaveOccurred())
					Expect(actual).To(Equal(&expected))

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("there are no students with classes in the DB", func() {
				It("should return an empty list", func() {
					expected := []model.StudentClass{}

					actual, err := studentRepo.FetchWithClass()
					Expect(err).NotTo(HaveOccurred())
					Expect(actual).To(Equal(&expected))

					err = db.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Describe("Class repository", func() {
			BeforeEach(func() {
				classes := []model.Class{
					{
						Name:       "Mathematics",
						Professor:  "Dr. Smith",
						RoomNumber: 101,
					},
					{
						Name:       "Physics",
						Professor:  "Dr. Johnson",
						RoomNumber: 102,
					},
					{
						Name:       "Chemistry",
						Professor:  "Dr. Lee",
						RoomNumber: 103,
					},
				}

				for _, c := range classes {
					err := conn.Create(&c).Error
					Expect(err).ShouldNot(HaveOccurred())
				}
			})

			When("there are classes in the database", func() {
				expectedClasses := []model.Class{
					{
						ID:         1,
						Name:       "Mathematics",
						Professor:  "Dr. Smith",
						RoomNumber: 101,
					},
					{
						ID:         2,
						Name:       "Physics",
						Professor:  "Dr. Johnson",
						RoomNumber: 102,
					},
					{
						ID:         3,
						Name:       "Chemistry",
						Professor:  "Dr. Lee",
						RoomNumber: 103,
					},
				}

				It("should return all classes", func() {
					classes, err := classRepo.FetchAll()
					Expect(err).To(BeNil())
					Expect(classes).To(HaveLen(3))
					Expect(classes).To(ConsistOf(expectedClasses))
				})
			})

			When("there are no classes in the database", func() {
				BeforeEach(func() {
					err = db.Reset(conn, "classes")
					Expect(err).ShouldNot(HaveOccurred())
				})

				It("should return an empty list of classes", func() {
					classes, err := classRepo.FetchAll()
					Expect(err).To(BeNil())
					Expect(classes).To(HaveLen(0))
				})
			})
		})
	})
})
