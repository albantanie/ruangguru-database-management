package main_test

import (
	main "a21hc3NpZ25tZW50"
	"fmt"
	"time"

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

	var sessionService service.SessionService

	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
		Schema:       "public",
	}
	conn, err := main.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	studentRepo = repo.NewStudentRepo(conn)
	userRepo = repo.NewUserRepo(conn)
	sessionRepo = repo.NewSessionRepo(conn)

	sessionService = service.NewSessionService(sessionRepo)

	BeforeEach(func() {
		// Drop table
		_, err = conn.Exec("DROP TABLE IF EXISTS students CASCADE")
		Expect(err).ShouldNot(HaveOccurred())

		_, err = conn.Exec("DROP TABLE IF EXISTS users CASCADE")
		Expect(err).ShouldNot(HaveOccurred())

		_, err = conn.Exec("DROP TABLE IF EXISTS sessions CASCADE")
		Expect(err).ShouldNot(HaveOccurred())

		// Create table
		err = main.SQLExecute(conn)
		Expect(err).ShouldNot(HaveOccurred())

		err = main.Reset(conn, "students")
		err = main.Reset(conn, "users")
		err = main.Reset(conn, "sessions")
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

						result, err := userRepo.FetchByID(1)
						Expect(result.Username).To(Equal(user.Username))
						Expect(result.Password).To(Equal(user.Password))

						err = main.Reset(conn, "users")
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

						result, err := userRepo.FetchByID(1)
						Expect(result.Username).To(Equal(user.Username))
						Expect(result.Password).To(Equal(user.Password))

						err = userRepo.CheckAvail(user)
						Expect(err).ShouldNot(HaveOccurred())

						err = main.Reset(conn, "users")
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

						result, err := sessionRepo.FetchByID(1)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						err = main.Reset(conn, "sessions")
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

						result, err := sessionRepo.FetchByID(1)
						Expect(err).ShouldNot(HaveOccurred())
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						err = sessionRepo.DeleteSession("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).ShouldNot(HaveOccurred())

						_, err = sessionRepo.FetchByID(1)
						Expect(err).To(Equal(fmt.Errorf("sql: no rows in result set")))

						err = main.Reset(conn, "sessions")
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

						result, err := sessionRepo.FetchByID(1)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						sessionUpdate := model.Session{
							Token:    "cc03dbac-4085-22ba-75fe-103f9a01b6d5",
							Username: "aditira",
							Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
						}
						err = sessionRepo.UpdateSessions(sessionUpdate)
						Expect(err).ShouldNot(HaveOccurred())

						result, err = sessionRepo.FetchByID(1)
						Expect(result.Token).To(Equal(sessionUpdate.Token))

						err = main.Reset(conn, "sessions")
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

						result, err := sessionRepo.FetchByID(1)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						tokenFound, err := sessionService.TokenValidity("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).ShouldNot(HaveOccurred())
						Expect(tokenFound.Token).To(Equal("cc03dbea-4085-47ba-86fe-020f5d01a9d8"))
						Expect(tokenFound.Username).To(Equal("aditira"))

						err = main.Reset(conn, "sessions")
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

						result, err := sessionRepo.FetchByID(1)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						tokenFound, err := sessionService.TokenValidity("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).To(Equal(fmt.Errorf("Token is Expired!")))
						Expect(tokenFound).To(Equal(model.Session{}))

						err = main.Reset(conn, "sessions")
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

						result, err := sessionRepo.FetchByID(1)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						err = sessionRepo.SessionAvailName("aditira")
						Expect(err).ShouldNot(HaveOccurred())

						err = main.Reset(conn, "sessions")
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

						result, err := sessionRepo.FetchByID(1)
						Expect(result.Token).To(Equal(session.Token))
						Expect(result.Username).To(Equal(session.Username))

						res, err := sessionRepo.SessionAvailToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
						Expect(err).ShouldNot(HaveOccurred())
						Expect(res.Token).To(Equal(session.Token))
						Expect(res.Username).To(Equal(session.Username))

						err = main.Reset(conn, "sessions")
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
						Class:   "A",
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())

					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("read all student data from students table database postgres", func() {
				It("should return a list of student data", func() {
					student1 := model.Student{
						Name:    "John",
						Address: "Jl. Raya",
						Class:   "A",
					}
					err := studentRepo.Store(&student1)
					Expect(err).ShouldNot(HaveOccurred())

					student2 := model.Student{
						Name:    "Doe",
						Address: "Jl. Melati",
						Class:   "B",
					}
					err = studentRepo.Store(&student2)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchAll()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result).To(HaveLen(2))
					Expect(result[0].Name).To(Equal(student1.Name))
					Expect(result[1].Name).To(Equal(student2.Name))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("adding new student data to students table in the database", func() {
				It("should save the new student data to students table in the database", func() {
					student := model.Student{
						Name:    "John",
						Address: "123 Main St",
						Class:   "Programming",
					}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("fetching all student data from students table in the database", func() {
				It("should return a list of all student data", func() {
					students := []model.Student{
						{Name: "John", Address: "123 Main St", Class: "Programming"},
						{Name: "Jane", Address: "456 Park Ave", Class: "Design"},
						{Name: "James", Address: "789 Broadway", Class: "Database"},
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
						Expect(result[i].Class).To(Equal(student.Class))
					}

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("fetching a single student data by id from students table in the database", func() {
				It("should return a single student data", func() {
					student := model.Student{Name: "John", Address: "123 Main St", Class: "Programming"}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(student.Name))
					Expect(result.Address).To(Equal(student.Address))
					Expect(result.Class).To(Equal(student.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("updating student data in students table in the database", func() {
				It("should update the existing student data in students table in the database", func() {
					student := model.Student{Name: "John", Address: "123 Main St", Class: "Programming"}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					newStudent := model.Student{Name: "Jane", Address: "456 Park Ave", Class: "Design"}
					err = studentRepo.Update(1, &newStudent)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).ShouldNot(HaveOccurred())
					Expect(result.Name).To(Equal(newStudent.Name))
					Expect(result.Address).To(Equal(newStudent.Address))
					Expect(result.Class).To(Equal(newStudent.Class))

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("deleting student data in students table in the database", func() {
				It("should delete the existing student data in students table in the database", func() {
					student := model.Student{Name: "John", Address: "123 Main St", Class: "Programming"}
					err := studentRepo.Store(&student)
					Expect(err).ShouldNot(HaveOccurred())

					err = studentRepo.Delete(1)
					Expect(err).ShouldNot(HaveOccurred())

					result, err := studentRepo.FetchByID(1)
					Expect(err).Should(HaveOccurred())
					Expect(result).To(BeNil())

					err = main.Reset(conn, "students")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})
})
